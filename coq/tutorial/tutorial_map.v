Require Import tutorial_lib.

Section Image.

Variable U V: Type.

Notation Domain   := (MySet U).
Notation Codomain := (MySet V).

(* Im C f ==> f(C) = { f(a) | a ∈ C } *)
Inductive Im (C: Domain) (f: U -> V) : Codomain :=
    Im_intro: forall x: U, x ∈ C -> forall y: V, y = f x -> y ∈ (Im C f).

(* x ∈ C ⇒ f(x) ∈ f(C) *)
Lemma Im_def: forall (C: Domain) (f: U -> V) (x: U),
    x ∈ C -> f x ∈ Im C f.
Proof.
intros.
apply Im_intro with x. trivial. trivial.
Qed.

(* y ∈ f(C) ⇔ ∃x [x ∈ C ∧ y = f(x)] *)
Lemma Im_elem: forall (C: Domain) (f: U -> V) (y: V),
    y ∈ Im C f <-> exists x: U, x ∈ C /\ y = f x.
Proof.
intros. split.
-
    intros.
    destruct H. exists x. split.
        trivial. trivial.
-
    intros.
    destruct H. destruct H.
    rewrite H0. apply Im_def. trivial.
Qed.

Lemma Im_subset: forall (A B: Domain) (f: U -> V),
    (A ⊆ B) -> Im A f ⊆ Im B f.
Proof.
intros.
unfold Included. intros.
destruct H0.
rewrite H1.
apply Im_def.
apply H.
trivial.
Qed.

Lemma Im_union: forall (A B: Domain) (f: U -> V),
    Im (A ∪ B) f = (Im A f ∪ Im B f).
Proof.
intros.
seteq.
-
    intro. intro.
    destruct H.
    rewrite H0.
    destruct H.
    --
        left.
        apply Im_def.
        trivial.
    --
        right.
        apply Im_def.
        trivial.
-
    intro y. intro.
    apply Im_elem.
    destruct H as [y|y].
    --
        apply Im_elem in H.
        destruct H. destruct H.
        exists x. split.
            left; trivial. 
            trivial.
    --
        apply Im_elem in H.
        destruct H. destruct H.
        exists x. split.
            right; trivial. trivial.
Qed.

Lemma Im_inter: forall (A B: Domain) (f: U -> V),
    Im (A ∩ B) f ⊆ (Im A f ∩ Im B f).
Proof.
intros.
    intro. intro. destruct H. destruct H.
    rewrite H0.
    split.
        - apply Im_def; trivial.
        - apply Im_def; trivial.
Qed.
        
(* InvIm D f ==> f⁻¹(D) = { a ∈ C | f(a) ∈ D } *)
Inductive InvIm (D: Codomain) (f: U -> V) : Domain :=
    InvIm_intro: forall x: U, f x ∈ D -> (x ∈ InvIm D f).

Lemma InvIm_def: forall (D: Codomain) (f: U -> V) (x: U),
    x ∈ InvIm D f <-> f x ∈ D.
Proof.
intros.
split.
-
    intros. destruct H. trivial.
-
    intros. apply InvIm_intro. trivial.
Qed.

Lemma InvImSet: forall (D: Codomain) (f: U -> V),
    InvIm D f = InvIm (Im (Full_set U) f ∩ D) f.
Proof.
intros.
seteq.
-
    intro. intro.
    apply InvIm_def in H.
    apply InvIm_def. split.
        -- apply Im_def. apply Full_intro.
        -- trivial.
-
    intro. intro.
    apply InvIm_def.
    apply InvIm_def in H.
    destruct H as [y].
    trivial.
Qed.

Lemma InvIm_subset: forall (C D: Codomain) (f: U -> V),
    (C ⊆ D) -> InvIm C f ⊆ InvIm D f.
Proof.
intros.
intro. intro.
apply InvIm_def.
apply InvIm_def in H0.
apply H in H0.
trivial.
Qed.

Lemma InvIm_union: forall (C D: Codomain) (f: U -> V),
    InvIm (C ∪ D) f = (InvIm C f ∪ InvIm D f).
Proof.
intros.
seteq.
-
    intro. intro.
    destruct H.
    remember (f x) as y.
    destruct H as [y|y].
    -- left. apply InvIm_def. rewrite <- Heqy. trivial.
    -- right. apply InvIm_def. rewrite <- Heqy. trivial.
-
    intro. intro.
    apply InvIm_def.
    destruct H.
    -- apply InvIm_def in H. left. trivial.
    -- apply InvIm_def in H. right. trivial.
Qed.

Lemma InvIm_inter: forall (C D: Codomain) (f: U -> V),
    InvIm (C ∪ D) f = (InvIm C f ∪ InvIm D f).
Proof.
intros.
seteq.
-
    intro. intro.
    apply InvIm_def in H.
    remember (f x) as y.
    destruct H as [y|y].
    -- left; apply InvIm_def; rewrite <- Heqy; trivial.
    -- right; apply InvIm_def; rewrite <- Heqy; trivial.
-
    intro. intro.
    apply InvIm_def.
    destruct H.
    -- left; apply InvIm_def in H; trivial.
    -- right; apply InvIm_def in H; trivial.
Qed.

Lemma InvIm_minus: forall (C D: Codomain) (f: U -> V),
    InvIm ((Full_set V) \ C) f = (Full_set U \ InvIm C f).
Proof.
intros.
seteq.
-
    intro. intro.
    apply InvIm_def in H.
    destruct H.
    split.
    -- apply Full_intro.
    -- intro; apply H0; apply InvIm_def in H1; contradiction.
-
    intro; intro.
    destruct H.
    apply InvIm_def.
    split.
    -- apply Full_intro.
    -- intro; apply H0; apply InvIm_def; trivial.
Qed.

Lemma Im_InvIm: forall (C: Codomain) (f: U -> V),
    Im (InvIm C f) f = (C ∩ Im (Full_set U) f).
Proof.
intros.
seteq.
-
    intro y; intro.
    apply Im_elem in H.
    destruct H. destruct H.
    apply InvIm_def in H.
    rewrite <- H0 in H.
    split.
    -- trivial.
    -- apply Im_elem; exists x; split.
        --- apply Full_intro.
        --- trivial.
-
    intro y; intro.
    destruct H as [y H0 H1].
    apply Im_elem in H1.
    destruct H1.
    destruct H.
    apply Im_elem.
    exists x.
    split.
    -- apply InvIm_def; rewrite <- H1; trivial.
    -- trivial.
Qed.

Lemma subset_Im_InvIm: forall (C D: Codomain) (f: U -> V),
    (C ∩ Im (Full_set U) f ⊆ D ∩ Im (Full_set U) f)
        -> InvIm C f ⊆ InvIm D f.
Proof.
intros.
intro; intro.
apply InvIm_def in H0.
remember (f x) as y.
unfold Included in H.
destruct H with y.
-
    split.
    --
        trivial.
    --
        apply Im_elem; exists x; split.
        --- apply Full_intro; trivial.
        --- trivial.
-
    apply InvIm_def.
    rewrite <- Heqy.
    trivial.
Qed.

End Image.
