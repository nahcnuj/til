(* http://herb.h.kobe-u.ac.jp/coq/coq.pdf *)

Require Import Classical.

Section Propositional_logic.

(* 命題 P, Q, R *)
Variable P Q R: Prop.

Lemma HA1: P -> (Q -> P).
Proof.
intro.   (* P を仮定に入れる *)
intro.   (* Q を仮定に入れる *)
trivial. (* 仮定より P は明らか *)
Qed.

Lemma Iand: P -> (Q -> P /\ Q).
Proof.
intros.  (* P, Q を仮定に入れる *)
split.   (* ゴールの論理積を分解 *)
trivial. (* P *)
trivial. (* Q *)
Qed.

Lemma And1: P /\ Q -> P.
Proof.
intro.      (* P /\ Q を仮定 *)
destruct H. (* 仮定 H: P /\ Q を分解 *)
trivial.    (* P *)
Qed.

Lemma Or1: P -> P \/ Q.
Proof.
intro.      (* P を仮定 *)
left.       (* ゴールの左の項 P を指定 *)
trivial.
Qed.

Lemma EMDN: P \/ ~P -> (~~P -> P).
Proof.
intro.
destruct H.
trivial.        (* P を仮定したときは明らかに P *)
contradiction.  (* ~P を仮定すると ~~P は矛盾 *)
Qed.

Lemma Cont: (P -> Q) -> (~Q -> ~P).
Proof.
intros.     (* ~P が残る．これは P -> False と同値 *)
intro.      (* P が仮定に入る *)
apply H0.   (* H0: ~Q すなわち Q -> False を適用．Q を示す *)
apply H.    (* H: P -> Q を適用．P を示す *)
trivial.    (* 仮定より P *)
Qed.

Lemma HA2: (P -> Q) -> ((P -> ~Q) -> ~P).
Proof.
intros.
intro.
(* apply H in H1.   (* H1: P を H: P -> Q で Q に書き換え *) *)
specialize(H H1).   (* H: P -> Q を H1: P で Q に書き換え *)
specialize(H0 H1).  (* H0: P -> ~Q を H1: P で ~Q に書き換え *)
contradiction.      (* 仮定に Q と ~Q があるので矛盾 *)
Qed.

Lemma Ex1_1: (P -> (Q -> R)) -> ((P -> Q) -> (P -> R)).
intros.
apply H.
trivial.
apply H0.
trivial.
Qed.

Lemma Ex1_2: (~(P /\ Q) -> (P -> ~Q)).
intros.
intro.
apply H.
split.
trivial.
trivial.
Qed.

Lemma Ex1_3: ((P /\ Q) -> R) -> (P -> (Q -> R)).
intros.
apply H.
split.
trivial.
trivial.
Qed.

Lemma Ex1_4: (Q -> P) -> ((R -> P) -> ((Q \/ R) -> P)).
intros.
destruct H1.
apply H.
trivial.
apply H0.
trivial.
Qed.

Lemma Ex1_5: (~P \/ ~Q) -> ~(P /\ Q).
intros.
intro.
destruct H.
destruct H0.
contradiction.
destruct H0.
contradiction.
Qed.

(*Require Import Classical.*)
Lemma Contra: (~Q -> ~P) -> (P -> Q).
intros.
apply NNPP.
intro.
specialize(H H1).
contradiction.
Qed.

Lemma Peirce: ((P -> Q) -> P) -> P.
Proof.
intros.
apply imply_to_or in H.  (* A -> B ≡ ~A \/ B *)
destruct H.
apply imply_to_and in H. (* ~(A -> B) ≡ A /\ ~B *)
destruct H.
trivial.
trivial.
Qed.

Lemma Ex2_1: ((P -> ~P) -> Q) -> ((P -> Q) -> Q).
intros.
apply NNPP.
intro.
apply imply_to_or in H.
destruct H.
apply imply_to_and in H.
destruct H.
specialize(H0 H).
contradiction.
contradiction.
Qed.

Lemma Ex2_2: (~Q -> ~P) -> ((~Q -> P) -> Q).
intros.
apply NNPP.
intro.
specialize(H H1).
specialize(H0 H1).
contradiction.
Qed.

Lemma Ex2_3: (P -> Q) \/ (Q -> P).
apply NNPP.
intro.
apply not_or_and in H.
destruct H.
apply not_imply_elim in H.
apply not_imply_elim2 in H0.
contradiction.
Qed.

End Propositional_logic.

Section Predicate_logic.

(* 対象領域 A *)
Variable A: Type.
(* A 上の 1 変数述語 P(x), Q(x) *)
Variable P Q: A -> Prop.
(* A 上の 2 変数述語 R(x, y) *)
Variable R: A -> (A -> Prop).
(* A の要素 t *)
Variable t: A.

(* (∀x P(x)) ⇒ P(t) *)
Lemma all_imply: (forall x: A, P x) -> P t.
Proof.
intro.
specialize(H t).
trivial.
Qed.

(* P(t) ⇒ (∃x P(x)) *)
Lemma imply_exists: P t -> (exists x: A, P x).
Proof.
intros.
exists t.
trivial.
Qed.

(* (∀x P(x)) ⇒ (∀y P(y)) *)
Lemma alpha_all: (forall x: A, P x) -> forall y: A, P y.
intro.
intro.
specialize(H y).
trivial.
Qed.

(* (∀x ¬P(x)) ⇒ ¬∃x P(x) *)
Lemma all_not_not_ex: (forall x: A, ~(P x)) -> (~exists x: A, P x).
intros.
intro.
destruct H0.
specialize(H x).
contradiction.
Qed.

Ltac ok := trivial; contradiction.

Lemma Ex4_1: ~(exists x, P x) -> forall x, ~P x.
intros.
intro.
destruct H.
exists x.
ok.
Qed.

Lemma Ex4_2: (exists x, ~P x) -> ~(forall x, P x).
intros.
intro.
destruct H.
specialize(H0 x).
ok.
Qed.

Lemma Ex4_3: (forall x, P x /\ Q x) -> (forall x, P x) /\ (forall x, Q x).
intros.
split.
intro.
    specialize(H x).
    destruct H.
ok.
intro.
    specialize(H x).
    destruct H.
ok.
Qed.

Lemma Ex4_4: (exists x, P x \/ Q x) -> (exists x, P x) \/ (exists x, Q x).
intros.
destruct H.
destruct H.
left.
    exists x.
ok.
right.
    exists x.
ok.
Qed.

(*Require Import Classical.*)

Lemma not_all_to_ex_not: ~(forall x, P x) -> exists x, ~P x.
Proof.
intros.
apply NNPP.
intro.
apply H.
intro; apply NNPP; intro.
apply H0.
exists x.
ok.
Qed.

Ltac 背理法で示す := apply NNPP; intro.

Lemma Ex5_1: (~ ~ exists x, P x) -> exists x, ~ ~ P x.
intros.
背理法で示す.
destruct H.
intro.
destruct H.
apply H0.
exists x.
intro.
ok.
Qed.

Lemma Ex5_2: (~ exists x, ~ P x) -> forall x, P x.
intros.
背理法で示す.
apply H.
exists x.
ok.
Qed.

End Predicate_logic.

Section Ensembles.

(*Require Import Classical.*)
Ltac ok := trivial; contradiction.
Ltac 背理法で示す := apply NNPP; intro.

Variable U: Type.
Definition MySet := U -> Prop.

Variable A B C D: MySet.

(* x ∈ A *)
Definition In (A: MySet) (x: U)  := A x.
(* A ⊆ B ⇔ ∀x∈U (x∈A ⇒ x∈B) *)
Definition Included (A B: MySet) := forall x: U, In A x -> In B x.

(* 全体集合Ω *)
Inductive Full_set : MySet := Full_intro: forall x: U, In Full_set x.
(* 空集合∅ *)
Inductive Empty_set : MySet := . (* なにもない *)

Definition Full_set2  : MySet := fun x: U => True.
Definition Empty_set2 : MySet := fun x: U => False.

(* x ∈ A∪B ⇔ x∈A ∨ x∈B *)
Inductive Union (A B: MySet) : MySet :=
    | Union_introl : forall x: U, In A x -> In (Union A B) x
    | Union_intror : forall x: U, In B x -> In (Union A B) x.

(* x ∈ A∩B ⇔ x∈A ∧ x∈B *)
Inductive Intersection (A B: MySet) : MySet :=
    | Intersection_intro: forall x: U, In A x -> In B x -> In (Intersection A B) x.

Notation "x ∈ A" := (In A x)           (at level 55, no associativity).
Notation "A ⊆ B" := (Included A B)     (at level 54, no associativity).
Notation "A ∩ B" := (Intersection A B) (at level 53, right associativity).
Notation "A ∪ B" := (Union A B)        (at level 53, right associativity).
Notation Ω       := (Full_set).
Notation Ø       := (Empty_set).

Lemma in_or_not : forall A, forall x, (x ∈ A) \/ ~(x ∈ A).
Proof.
intros.
apply classic.
Qed.

Lemma subset_transitive : (A ⊆ B) /\ (B ⊆ C) -> (A ⊆ C).
Proof.
unfold Included.
intros.
destruct H.
apply H1.
apply H.
ok.
Qed.

Ltac unfolds := unfold Included; intros.

Lemma empty_subset : forall A: MySet, Ø ⊆ A.
Proof.
intros.
unfolds.
destruct H.
Qed.

Lemma subset_full : forall A: MySet, A ⊆ Ω.
Proof.
unfolds.
apply Full_intro.
Qed.

Definition Same_set (A B: MySet) := A ⊆ B /\ B ⊆ A.
Axiom Extensionality_Sets : forall A B: MySet, Same_set A B -> A = B.
Ltac seteq := apply Extensionality_Sets; unfold Same_set; split.

Lemma union_id : (A ∪ A) = A.
Proof.
seteq.
unfolds.
destruct H.
ok.
ok.
unfolds.
apply Union_introl.
ok.
Qed.

Lemma union_comm : (A ∪ B) = (B ∪ A).
Proof.
seteq.
unfolds.
destruct H.
apply Union_intror. ok.
apply Union_introl. ok.
unfolds.
destruct H.
apply Union_intror. ok.
apply Union_introl. ok.
Qed.

Lemma union_assoc : (A ∪ (B ∪ C)) = ((A ∪ B) ∪ C).
Proof.
seteq.
unfolds.
    destruct H.
        apply Union_introl. apply Union_introl. ok.
    destruct H.
        apply Union_introl. apply Union_intror. ok.
    apply Union_intror. ok.
unfolds.
    destruct H.
        destruct H.
            apply Union_introl. ok.
        apply Union_intror, Union_introl. ok.
        apply Union_intror, Union_intror. ok.
Qed.

Lemma union_subset : A ⊆ (A ∪ B) /\ B ⊆ (A ∪ B).
Proof.
split.
    unfolds.
        apply Union_introl. ok.
    unfolds. 
        apply Union_intror. ok.
Qed.

Lemma subset_union : A ⊆ C /\ B ⊆ C -> A ∪ B ⊆ C.
Proof.
intros.
destruct H.
unfolds.
destruct H1.
    apply H. ok.
    apply H0. ok.
Qed.

Lemma subset_union_absorp : A ⊆ B <-> (A ∪ B) = B.
Proof.
split.
    intro. seteq.
        unfolds. destruct H0.
            apply H. ok.
            ok.
        unfolds. apply Union_intror. ok.
    intros. rewrite <- H. apply union_subset.
Qed.

Lemma intersec_id : (A ∩ A) = A.
Proof.
seteq.
    unfolds. destruct H. ok.
    unfolds. apply Intersection_intro. ok. ok.
Qed.

Lemma intersec_comm : (A ∩ B) = (B ∩ A).
Proof.
seteq.
    unfolds. destruct H.
        apply Intersection_intro. ok. ok.
    unfolds. destruct H.
        apply Intersection_intro. ok. ok.
Qed.

Lemma intersec_assoc : (A ∩ (B ∩ C)) = ((A ∩ B) ∩ C).
Proof.
seteq.
    unfolds.
        destruct H. destruct H0.
        apply Intersection_intro. apply Intersection_intro.
        ok. ok. ok.
    unfolds.
        destruct H. destruct H.
        apply Intersection_intro. ok.
        apply Intersection_intro. ok. ok.
Qed.

Lemma intersec_subset : (A ∩ B) ⊆ A /\ (A ∩ B) ⊆ B.
Proof.
split.
    unfolds. destruct H. ok.
    unfolds. destruct H. ok.
Qed.

Lemma subset_intersec : (C ⊆ A) /\ (C ⊆ B) -> (C ⊆ (A ∩ B)).
Proof.
intro. destruct H. unfolds. split.
    apply H. ok.
    apply H0. ok.
Qed.

Lemma subset_intersec_absorp : (A ⊆ B) <-> (A ∩ B) = A.
Proof.
split.
    intro. seteq.
        unfolds. destruct H0. ok.
        unfolds. split. ok. apply H in H0. ok.
    intro. rewrite <- H. unfolds. destruct H0. ok.
Qed.

Inductive Disjoint (A B: MySet) : Prop :=
    | Disjoint_intro : (forall x: U, ~(x ∈ (A ∩ B))) -> Disjoint A B.

Lemma disjoint_empty : Disjoint A B <-> (A ∩ B) = Ø.
Proof.
split.
    intros. seteq.
        unfolds. destruct H. specialize (H x). destruct H.
            destruct H0. split. ok. ok.
        unfolds. ok.
    intros. apply Disjoint_intro. intro. rewrite H. intro. ok.
Qed.

Lemma intersec_of_union : (A ∪ (B ∩ C)) = ((A ∪ B) ∩ (A ∪ C)).
Proof.
seteq.
    unfolds. destruct H.
        split.
            apply Union_introl. ok.
            apply Union_introl. ok.
        split.
            destruct H. apply Union_intror. ok.
            destruct H. apply Union_intror. ok.
    unfolds. destruct H. 
        destruct H. apply Union_introl. ok.
        destruct H0. apply Union_introl. ok.
        apply Union_intror. split. ok. ok.
Qed.

Lemma union_of_intersec : (A ∩ (B ∪ C)) = ((A ∩ B) ∪ (A ∩ C)).
Proof.
seteq.
    unfolds. destruct H. destruct H0.
        left. apply Intersection_intro. ok. ok.
        right. apply Intersection_intro. ok. ok.
    unfolds. destruct H. destruct H. apply Intersection_intro. ok. left. ok.
    split.
        destruct H. ok.
        destruct H. right. ok.
Qed.

Lemma absorption_1 : (A ∪ (A ∩ B)) = A.
Proof.
seteq.
    unfolds.
        destruct H. ok.
        destruct H. ok.
    unfolds.
        left. ok.
Qed.

Lemma absorption_2 : (A ∩ (A ∪ B)) = A.
Proof.
seteq.
    unfolds.
        destruct H. ok.
    unfolds. split.
        ok.
        left. ok.
Qed.

Lemma subset_seq : (A ⊆ B /\ B ⊆ C) <-> (A ∪ B = B ∩ C).
Proof.
split.
    intros. destruct H. apply subset_union_absorp in H. seteq.
        intro. intro. rewrite H in H1. split.
            ok.
            apply H0 in H1. ok.
        unfolds. destruct H1. right. ok.
    intros. split.
        unfolds. assert (x ∈ (A ∪ B)). left. ok. rewrite H in H1. destruct H1. ok.
        unfolds. assert (x ∈ (A ∪ B)). right. ok. rewrite H in H1. destruct H1. ok.
Qed.

Definition Setminus (A B: MySet) : MySet :=
    fun x: U => (x ∈ A) /\ ~(x ∈ B).
Notation "A \ B" := (Setminus A B) (at level 60, no associativity).

Lemma setminus : forall A B, forall x,
    (x ∈ A) -> (~(x ∈ B) -> x ∈ (A \ B)).
Proof.
intros. unfold In. unfold Setminus. split. ok. ok.
Qed.

Lemma setminus_subset : (A \ B) ⊆ A.
Proof.
intro. intro. destruct H. ok.
Qed.

Lemma setminus_intersec_empty : (A \ B) ∩ B = Ø.
Proof.
seteq.
    unfolds. destruct H. destruct H. ok.
    unfolds. ok.
Qed.

Lemma union_setminus_intersec : (A \ B) ∪ (A ∩ B) = A.
Proof.
seteq.
    unfolds. destruct H.
        destruct H. ok.
        destruct H. ok.
    unfolds. destruct (in_or_not B x).
        right. split. ok. ok.
        left. unfold In, Setminus. split. ok. ok.
Qed.

Lemma A_9 : (A ⊆ C /\ D ⊆ B) -> (A \ B) ⊆ (C \ D).
Proof.
intro.
destruct H.
unfolds. unfold In, Setminus in H1. destruct H1.
unfold In, Setminus. split.
    apply H in H1. ok.
    intro. apply H2. apply H0 in H3. ok.
Qed.

Lemma A_10 : (A ⊆ B) <-> (A \ B = Ø).
Proof.
split.
    intro. seteq.
        intro. intro. unfold In, Setminus in H0.
            destruct H0. apply H in H0. ok.
        intro. ok.
    intro. unfolds. destruct (in_or_not B x).
        ok.
        assert (x ∈ (A \ B)). unfold In, Setminus. split.
            ok. ok.
            rewrite H in H2. ok.
Qed.

Lemma A_11 : (A \ (A \ B)) = A ∩ B.
seteq.
    intro. intro. destruct H.
    unfold In, Setminus in H0. apply not_and_or in H0. destruct H0.
        split.
            ok.
            ok.
        split.
            ok.
            apply NNPP. ok.
    unfolds. destruct H. split.
        ok.
        unfold In, Setminus. intro. destruct H1. ok.
Qed.

Lemma A_12 : (A ⊆ C /\ B ⊆ C) -> (A ∩ B = Ø <-> A ⊆ (C \ B)).
Proof.
intro. destruct H. split.
    intro. apply disjoint_empty in H1. destruct H1.
        intro. intro. unfold In, Setminus. split.
            apply H in H2. ok.
            destruct (in_or_not B x).
                specialize (H1 x). destruct H1. split. ok. ok. ok.
    intro. seteq.
        intro. intro. destruct H2. apply H1 in H2. unfold In, Setminus in H2. destruct H2. ok.
        intro. intro. ok.
Qed.

Lemma A_13 : (A \ C) ⊆ ((A \ B) ∪ (B \ C)).
Proof.
intro. intro. unfold In, Setminus in H. destruct H.
destruct (in_or_not B x).
    right. unfold In, Setminus. split. ok. ok.
    left. unfold In, Setminus. split. ok. ok.
Qed.

Lemma A_14_1 : A \ (B ∪ C) = (A \ B) ∩ (A \ C).
Proof.
seteq.
    intro. intro. unfold In, Setminus in H. destruct H.
        assert (~ x ∈ B). intro. apply H0. left. ok.
        assert (~ x ∈ C). intro. apply H0. right. ok.
        split.
            split. ok. ok.
            split. ok. ok.
    intro. intro. destruct H.
        unfold In, Setminus in H. destruct H.
        unfold In, Setminus in H0. destruct H0.
        unfold In, Setminus. split. ok.
        intro. destruct H3. apply H1. ok. ok.
Qed.

Lemma A_14_2 : A \ (B ∩ C) = (A \ B) ∪ (A \ C).
Proof.
seteq.
    intro. intro. unfold In, Setminus in H. destruct H.
        destruct (in_or_not C x).
            left. unfold In, Setminus. split. ok. intro. apply H0. split. ok. ok.
            right. unfold In, Setminus. split. ok. ok.
    intro. intro. split.
        destruct H.
            unfold In, Setminus in H. destruct H. ok.
            unfold In, Setminus in H. destruct H. ok.
        intro. destruct H0. destruct H.
            unfold In, Setminus in H. destruct H. ok.
            unfold In, Setminus in H. destruct H. ok.
Qed.

Lemma A_15_1: (A ∪ B) \ C = (A \ C) ∪ (B \ C).
Proof.
seteq.
    intro. intro. destruct H. destruct H.
        left. split. ok. ok.
        right. split. ok. ok.
    intro. intro. destruct H.
        destruct H. split.
            left. ok.
            ok.
        destruct H. split.
            right. ok.
            ok.
Qed.

Lemma A_15_2: (A ∩ B) \ C = (A \ C) ∩ (B \ C).
Proof.
seteq.
    intro. intro. destruct H. destruct H. split.
        split. ok. ok.
        split. ok. ok.
    intro. intro. destruct H. destruct H. destruct H0. split. split. ok. ok. ok.
Qed.

Lemma A_15_3: (A \ B) ∩ C = (A ∩ C) \ B.
Proof.
seteq.
    intro. intro.
    destruct H. destruct H.
    split.
        split. ok. ok.
        ok.
    intro. intro.
    destruct H. destruct H. simpl in H.
    split.
        split. ok. ok.
        ok.
Qed.

Lemma A_15_4: ((A ∪ C) \ B) ⊆ (A \ B) ∪ C.
Proof.
intro. intro.
destruct H. destruct H.
    left. split. ok. ok.
    right. ok.
Qed.

End Ensembles.

Section Family.

Variable U Λ: Type.

Notation S := (MySet U).
Notation "x ∈ A" := (In U A x)           (at level 55, no associativity).
Notation "A ⊆ B" := (Included U A B)     (at level 54, no associativity).
Notation "A ∩ B" := (Intersection U A B) (at level 53, right associativity).
Notation "A ∪ B" := (Union U A B)        (at level 53, right associativity).
Notation Ω       := (Full_set U).
Notation Ø       := (Empty_set U).

Definition Family {Λ: _} {S: _} := Λ -> S.

Inductive UnionF (ℱ: Family) : S :=
    | unionf_intro: forall x: U, (exists λ: Λ, (x ∈ (ℱ λ))) -> x ∈ (UnionF ℱ).

Inductive InterF (ℱ: Family) : S :=
    | interf_intro: forall x: U, (forall λ: Λ, (x ∈ (ℱ λ))) -> x ∈ (InterF ℱ).

Lemma mem_unionf: forall ℱ: Family, forall λ0: Λ,
    ℱ λ0 ⊆ UnionF ℱ.
Proof.
intros. intro. intro. apply unionf_intro. exists λ0. trivial.
Qed.

Lemma mem_interf: forall ℱ: Family, forall λ0: Λ,
    InterF ℱ ⊆ ℱ λ0.
Proof.
intros. intro. intro. destruct H. apply H.
Qed.

Lemma unionf_inc: forall ℱ 𝒢: Family,
    (forall λ: Λ, ℱ λ ⊆ 𝒢 λ) -> UnionF ℱ ⊆ UnionF 𝒢.
Proof.
intros.
intro. intro.
unfold Included in H.
destruct H0. destruct H0 as [λ].
apply unionf_intro. exists λ.
apply H.
trivial.
Qed.

Lemma interf_inc: forall ℱ 𝒢: Family,
    (forall λ: Λ, ℱ λ ⊆ 𝒢 λ) -> InterF ℱ ⊆ InterF 𝒢.
Proof.
intros. intro. intro.
split. intro. apply H.
destruct H0. specialize (H0 λ).
trivial.
Qed.

End Family.

Notation S := (MySet _).
Notation "x ∈ A" := (In _ A x)           (at level 50, no associativity).
Notation "A ⊆ B" := (Included _ A B)     (at level 100, no associativity).
Notation "A ∩ B" := (Intersection _ A B) (at level 80, right associativity).
Notation "A ∪ B" := (Union _ A B)        (at level 80, right associativity).
Notation "A \ B" := (Setminus _ A B)       (at level 80, no associativity).

Ltac seteq := apply Extensionality_Sets; unfold Same_set; split.
