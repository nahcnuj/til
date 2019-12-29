Require Import tutorial_lib.

Section Relation.

Variable U: Type.
Notation RelOp := (U -> U -> Prop).

(* 反射律: ∀a∈A [a R a]*)
Definition Reflexive (R: RelOp) :=
    forall a: U, R a a.

(* 対称律: ∀a,b∈A [a R b ⇒ b R a] *)
Definition Symmetric (R: RelOp) :=
    forall a b: U, R a b -> R b a.

(* 推移律: ∀a,b,c∈A [a R b ∧ b R c ⇒ a R c] *)
Definition Transitive (R: RelOp) :=
    forall a b c: U, R a b -> R b c -> R a c.

(* 同値関係: 反射律・対称律・推移律を満たす 2 項関係 *)
Definition Equivalence_Relation (R: RelOp) :=
    Reflexive R /\ Symmetric R /\ Transitive R.

(* R を同値関係と仮定する *)
Variable R: RelOp.
Hypothesis equiv: Equivalence_Relation R.

Ltac unfold_eqrel :=
    destruct equiv as [?R [?S ?T]].

(* 同値類 [x] = { a ∈ A | x R a } *)
Inductive Eqclass (x: U) : MySet U :=
    | Eqclass_intro: forall a: U, R x a -> a ∈ (Eqclass x).
Notation "[ x ]" := (Eqclass x) (at level 30, no associativity).


Lemma Eqclass_id: forall x: U,
    x ∈ [x].
Proof.
intros.
unfold_eqrel.
apply Eqclass_intro.
apply R0.
Qed.

Lemma Eqclass_subset: forall x y: U,
    R x y -> [x] ⊆ [y].
Proof.
intros.
intro a; intro.
apply Eqclass_intro.
destruct H0.
unfold_eqrel.
apply (T y x a).
- apply S; trivial.
- trivial.
Qed.

Lemma Eqclass_eq: forall a a': U,
    R a a' -> [a] = [a'].
Proof.
intros.
unfold_eqrel.
seteq; split.
-
    destruct H0 as [x].
    apply (T a' a x).
    -- apply S; trivial.
    -- trivial.
-
    destruct H0 as [x].
    apply (T a a' x).
    -- trivial.
    -- trivial.
Qed.

Lemma Eqclass_neq: forall a a': U,
    ~ R a a' -> ([a] ∩ [a']) = Empty_set U.
Proof.
intros.
unfold_eqrel.
seteq.
-
    intro; intro.
    destruct H0.
    destruct H0.
    destruct H1 as [x].
    assert (R a a').
    -- apply (T a x a').
        --- trivial.
        --- apply S; trivial.
    -- contradiction.
-
    intro; intro.
    destruct H0.
Qed.

Definition EqFam := fun x: U => [x].

Variable t: U.
Eval compute in (UnionF _ _ EqFam).

Lemma Eqclass_all:
    UnionF _ _ EqFam = Full_set _.
Proof.
unfold_eqrel.
seteq.
split.
-
    intro; intro.
    apply unionf_intro.
    exists x.
    apply Eqclass_id.
Qed.

End Relation.
